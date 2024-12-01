<?php
declare(strict_types=1);

namespace App\Command\Day1;

use App\Service\FileReaderService;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand(name: 'aoc:1:1', description: 'Challenge Day 1 - Part 1')]
class PartOneCommand extends Command
{
    public function __construct(
        private readonly FileReaderService $fileReaderService,
        protected readonly ?string $name = null,
    ) {
        parent::__construct($name);
    }

    protected function configure(): void
    {
        $this
            ->addOption(
                'input',
                'i',
                InputOption::VALUE_REQUIRED,
                'File name to use as input',
                'example'
            );
    }

    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        /** @var string $inputFile */
        $inputFile = $input->getOption('input');

        $lines = $this->fileReaderService->readLines(1, $inputFile . '.txt');

        $listOne = [];
        $listTwo = [];

        foreach ($lines as $line) {
            $values = explode('   ', $line);

            $listOne[] = (int) $values[0];
            $listTwo[] = (int) $values[1];
        }

        sort($listOne);
        sort($listTwo);

        $total = 0;
        for ($i = 0; $i < count($listOne); $i++) {
            $total += self::getDifference($listOne[$i], $listTwo[$i]);
        }

        $output->writeln('Result: ' . $total);

        return Command::SUCCESS;
    }

    private static function getDifference(int $number1, int $number2): int
    {
        if ($number1 > $number2) {
            return $number1 - $number2;
        }

        return $number2 - $number1;
    }
}
