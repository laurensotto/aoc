<?php
declare(strict_types=1);

namespace App\Command\Day2;

use App\Service\FileReaderService;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Input\InputOption;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand(name: 'aoc:2:2', description: 'Challenge Day 2 - Part 2')]
class PartTwoCommand extends Command
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

        $lines = $this->fileReaderService->readLines(2, $inputFile . '.txt');

        $points = 0;

        foreach ($lines as $line) {
            $values   = str_split($line);
            $opponent = $values[0];
            $you      = $values[2];

            $points += self::getGameOutcome($opponent, $you);
        }

        $output->writeln('Winning score: ' . $points);

        return Command::SUCCESS;
    }
    private static function getGameOutCome(string $opponent, string $you): int
    {
        return match ($you) {
            'X' => self::getGameOutcomeLose($opponent),
            'Y' => self::getGameOutcomeDraw($opponent),
            'Z' => self::getGameOutcomeWin($opponent),
            default => throw new \Exception('Unknown choice: ' . $you),
        };
    }

    private static function getGameOutcomeWin(string $opponent): int
    {
        return match ($opponent) {
            'A' => 8,
            'B' => 9,
            'C' => 7,
            default => throw new \Exception('Unknown opponent choice: ' . $opponent),
        };
    }

    private static function getGameOutcomeDraw(string $opponent): int
    {
        return match ($opponent) {
            'A' => 4,
            'B' => 5,
            'C' => 6,
            default => throw new \Exception('Unknown opponent choice: ' . $opponent),
        };
    }

    private static function getGameOutcomeLose(string $opponent): int
    {
        return match ($opponent) {
            'A' => 3,
            'B' => 1,
            'C' => 2,
            default => throw new \Exception('Unknown opponent choice: ' . $opponent),
        };
    }
}
